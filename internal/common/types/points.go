package types

type StatPoint int16 // TODO change type?

const (
	POINT_NONE        StatPoint = iota // 0
	POINT_LEVEL                        // 1
	POINT_VOICE                        // 2
	POINT_EXP                          // 3
	POINT_NEXT_EXP                     // 4
	POINT_HP                           // 5
	POINT_MAX_HP                       // 6
	POINT_SP                           // 7
	POINT_MAX_SP                       // 8
	POINT_STAMINA                      // 9  ���׹̳�
	POINT_MAX_STAMINA                  // 10 �ִ� ���׹̳�

	POINT_GOLD             // 11
	POINT_ST               // 12 �ٷ�
	POINT_HT               // 13 ü��
	POINT_DX               // 14 ��ø��
	POINT_IQ               // 15 ���ŷ�
	POINT_DEF_GRADE        // 16 ...
	POINT_ATT_SPEED        // 17 ���ݼӵ�
	POINT_ATT_GRADE        // 18 ���ݷ� MAX
	POINT_MOV_SPEED        // 19 �̵��ӵ�
	POINT_CLIENT_DEF_GRADE // 20 �����
	POINT_CASTING_SPEED    // 21 �ֹ��ӵ� (��ٿ�Ÿ��*100) / (100 + �̰�) = ���� ��ٿ� Ÿ��
	POINT_MAGIC_ATT_GRADE  // 22 �������ݷ�
	POINT_MAGIC_DEF_GRADE  // 23 ��������
	POINT_EMPIRE_POINT     // 24 ��������
	POINT_LEVEL_STEP       // 25 �� ���������� �ܰ�.. (1 2 3 �� �� ���� 4 �Ǹ� ���� ��)
	POINT_STAT             // 26 �ɷ�ġ �ø� �� �ִ� ����
	POINT_SUB_SKILL        // 27 ���� ��ų ����Ʈ
	POINT_SKILL            // 28 ��Ƽ�� ��ų ����Ʈ
	POINT_WEAPON_MIN       // 29 ���� �ּ� ������
	POINT_WEAPON_MAX       // 30 ���� �ִ� ������
	POINT_PLAYTIME         // 31 �÷��̽ð�
	POINT_HP_REGEN         // 32 HP ȸ����
	POINT_SP_REGEN         // 33 SP ȸ����

	POINT_BOW_DISTANCE // 34 Ȱ �����Ÿ� ����ġ (meter)

	POINT_HP_RECOVERY // 35 ü�� ȸ�� ������
	POINT_SP_RECOVERY // 36 ���ŷ� ȸ�� ������

	POINT_POISON_PCT    // 37 �� Ȯ��
	POINT_STUN_PCT      // 38 ���� Ȯ��
	POINT_SLOW_PCT      // 39 ���ο� Ȯ��
	POINT_CRITICAL_PCT  // 40 ũ��Ƽ�� Ȯ��
	POINT_PENETRATE_PCT // 41 ����Ÿ�� Ȯ��
	POINT_CURSE_PCT     // 42 ���� Ȯ��

	POINT_ATTBONUS_HUMAN    // 43 �ΰ����� ����
	POINT_ATTBONUS_ANIMAL   // 44 �������� ������ % ����
	POINT_ATTBONUS_ORC      // 45 ���Ϳ��� ������ % ����
	POINT_ATTBONUS_MILGYO   // 46 �б����� ������ % ����
	POINT_ATTBONUS_UNDEAD   // 47 ��ü���� ������ % ����
	POINT_ATTBONUS_DEVIL    // 48 ����(�Ǹ�)���� ������ % ����
	POINT_ATTBONUS_INSECT   // 49 ������
	POINT_ATTBONUS_FIRE     // 50 ȭ����
	POINT_ATTBONUS_ICE      // 51 ������
	POINT_ATTBONUS_DESERT   // 52 �縷��
	POINT_ATTBONUS_MONSTER  // 53 ��� ���Ϳ��� ����
	POINT_ATTBONUS_WARRIOR  // 54 ���翡�� ����
	POINT_ATTBONUS_ASSASSIN // 55 �ڰ����� ����
	POINT_ATTBONUS_SURA     // 56 ���󿡰� ����
	POINT_ATTBONUS_SHAMAN   // 57 ���翡�� ����
	POINT_ATTBONUS_TREE     // 58 �������� ���� 20050729.myevan UNUSED5

	POINT_RESIST_WARRIOR  // 59 ���翡�� ����
	POINT_RESIST_ASSASSIN // 60 �ڰ����� ����
	POINT_RESIST_SURA     // 61 ���󿡰� ����
	POINT_RESIST_SHAMAN   // 62 ���翡�� ����

	POINT_STEAL_HP // 63 ������ ����
	POINT_STEAL_SP // 64 ���ŷ� ����

	POINT_MANA_BURN_PCT // 65 ���� ��

	/// ���ؽ� ���ʽ� ///

	POINT_DAMAGE_SP_RECOVER // 66 ���ݴ��� �� ���ŷ� ȸ�� Ȯ��

	POINT_BLOCK // 67 ������
	POINT_DODGE // 68 ȸ����

	POINT_RESIST_SWORD   // 69
	POINT_RESIST_TWOHAND // 70
	POINT_RESIST_DAGGER  // 71
	POINT_RESIST_BELL    // 72
	POINT_RESIST_FAN     // 73
	POINT_RESIST_BOW     // 74  ȭ��   ����   : ����� ����
	POINT_RESIST_FIRE    // 75  ȭ��   ����   : ȭ�����ݿ� ���� ����� ����
	POINT_RESIST_ELEC    // 76  ����   ����   : ������ݿ� ���� ����� ����
	POINT_RESIST_MAGIC   // 77  ����   ����   : �������� ���� ����� ����
	POINT_RESIST_WIND    // 78  �ٶ�   ����   : �ٶ����ݿ� ���� ����� ����

	POINT_REFLECT_MELEE // 79 ���� �ݻ�

	/// Ư�� ���ؽ� ///
	POINT_REFLECT_CURSE // 80 ���� �ݻ�
	POINT_POISON_REDUCE // 81 �������� ����

	/// �� �Ҹ�� ///
	POINT_KILL_SP_RECOVER   // 82 �� �Ҹ�� MP ȸ��
	POINT_EXP_DOUBLE_BONUS  // 83
	POINT_GOLD_DOUBLE_BONUS // 84
	POINT_ITEM_DROP_BONUS   // 85

	/// ȸ�� ���� ///
	POINT_POTION_BONUS     // 86
	POINT_KILL_HP_RECOVERY // 87

	POINT_IMMUNE_STUN // 88
	POINT_IMMUNE_SLOW // 89
	POINT_IMMUNE_FALL // 90
	//////////////////

	POINT_PARTY_ATTACKER_BONUS // 91
	POINT_PARTY_TANKER_BONUS   // 92

	POINT_ATT_BONUS // 93
	POINT_DEF_BONUS // 94

	POINT_ATT_GRADE_BONUS       // 95
	POINT_DEF_GRADE_BONUS       // 96
	POINT_MAGIC_ATT_GRADE_BONUS // 97
	POINT_MAGIC_DEF_GRADE_BONUS // 98

	POINT_RESIST_NORMAL_DAMAGE // 99

	POINT_HIT_HP_RECOVERY // 100
	POINT_HIT_SP_RECOVERY // 101
	POINT_MANASHIELD      // 102 ��ż�ȣ ��ų�� ���� �������� ȿ�� ����

	POINT_PARTY_BUFFER_BONUS       // 103
	POINT_PARTY_SKILL_MASTER_BONUS // 104

	POINT_HP_RECOVER_CONTINUE // 105
	POINT_SP_RECOVER_CONTINUE // 106

	POINT_STEAL_GOLD // 107
	POINT_POLYMORPH  // 108 ������ ���� ��ȣ
	POINT_MOUNT      // 109 Ÿ���ִ� ���� ��ȣ

	POINT_PARTY_HASTE_BONUS    // 110
	POINT_PARTY_DEFENDER_BONUS // 111
	POINT_STAT_RESET_COUNT     // 112 ���� �ܾ� ����� ���� ���� ���� ����Ʈ (1�� 1����Ʈ ���°���)

	POINT_HORSE_SKILL // 113

	POINT_MALL_ATTBONUS  // 114 ���ݷ� +x%
	POINT_MALL_DEFBONUS  // 115 ���� +x%
	POINT_MALL_EXPBONUS  // 116 ����ġ +x%
	POINT_MALL_ITEMBONUS // 117 ������ ����� x/10��
	POINT_MALL_GOLDBONUS // 118 �� ����� x/10��

	POINT_MAX_HP_PCT // 119 �ִ������ +x%
	POINT_MAX_SP_PCT // 120 �ִ����ŷ� +x%

	POINT_SKILL_DAMAGE_BONUS      // 121 ��ų ������ *(100+x)%
	POINT_NORMAL_HIT_DAMAGE_BONUS // 122 ��Ÿ ������ *(100+x)%

	// DEFEND_BONUS_ATTRIBUTES
	POINT_SKILL_DEFEND_BONUS      // 123 ��ų ��� ������
	POINT_NORMAL_HIT_DEFEND_BONUS // 124 ��Ÿ ��� ������
	// END_OF_DEFEND_BONUS_ATTRIBUTES

	// PC_BANG_ITEM_ADD
	POINT_PC_BANG_EXP_BONUS  // 125 PC�� ���� ����ġ ���ʽ�
	POINT_PC_BANG_DROP_BONUS // 126 PC�� ���� ��ӷ� ���ʽ�
	// END_PC_BANG_ITEM_ADD
	POINT_RAMADAN_CANDY_BONUS_EXP // �󸶴� ���� ����ġ ������

	POINT_ENERGY = 128 // 128 ���

	// ��� ui ��.
	// �������� ���� �ʱ⸸ Ŭ���̾�Ʈ���� ����� �� �ð��� POINT�� �����ϱ� ������ �̷��� �Ѵ�.
	// �� �β�����
	POINT_ENERGY_END_TIME = 129 // 129 ��� ���� �ð�

	POINT_COSTUME_ATTR_BONUS        = 130
	POINT_MAGIC_ATT_BONUS_PER       = 131
	POINT_MELEE_MAGIC_ATT_BONUS_PER = 132

	// �߰� �Ӽ� ����
	POINT_RESIST_ICE   = 133 //   �ñ� ����   : �������ݿ� ���� ����� ����
	POINT_RESIST_EARTH = 134 //   ���� ����   : �������ݿ� ���� ����� ����
	POINT_RESIST_DARK  = 135 //   ��� ����   : �������ݿ� ���� ����� ����

	POINT_RESIST_CRITICAL  = 136 // ũ��Ƽ�� ����	: ����� ũ��Ƽ�� Ȯ���� ����
	POINT_RESIST_PENETRATE = 137 // ����Ÿ�� ����	: ����� ����Ÿ�� Ȯ���� ����

	//POINT_MAX_NUM = 129	common/length.h
)
